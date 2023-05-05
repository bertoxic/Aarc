package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bertoxic/aarc/config"
	"github.com/bertoxic/aarc/drivers"
	"github.com/bertoxic/aarc/forms"
	"github.com/bertoxic/aarc/helpers"
	"github.com/bertoxic/aarc/models"
	"github.com/bertoxic/aarc/render"
	"github.com/bertoxic/aarc/repository"
	"github.com/bertoxic/aarc/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(db *drivers.DB, app *config.AppConfig) *Repository{
    return &Repository{
        App: app,
        DB: dbrepo.NewPostgresDBRepo(db.SQL, app),
    }
}

func NewHandlers(r *Repository) {
	Repo = r
}

type JsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
	Name string `json:"name"`
	Email string `json:"email"`
	Pass string `json:"pass"`
}

func (m *Repository) Register (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	err :=r.ParseForm()
	if err != nil {
		resp := JsonResponse{
			OK:      false,
			Message: "Internal server error",
		}
		out, err := json.MarshalIndent(resp, "", "   ")
		if err != nil {
			log.Println(err)
		}
		
		w.Write(out)
		
		return
	}

	name:=r.Form.Get("name")
	email:=r.Form.Get("email")
	pass:=r.Form.Get("password")
	resp := JsonResponse{
		OK:      false,
		Message: " server running nice",
		Name: name,
		Email: email,
		Pass: pass,
	}

	isEmail, err := m.DB.IsEmailUsed(email)
	if err != nil {
		resp := JsonResponse{
			OK:      false,
			Message: "Internal server error",
		}
		out, err := json.MarshalIndent(resp, "", "   ")
		if err != nil {
			log.Println(err)
		}
		//w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		
		return
	}

	if isEmail {
		
			resp := JsonResponse{
				OK:      false,
				Message: "Email already used",
			}
			out, err := json.MarshalIndent(resp, "", "   ")
			if err != nil {
				log.Println(err)
			}
			//w.Header().Set("Content-Type", "application/json")
			w.Write(out)
			
			return

}

	out, _ := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}
//	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	verificationCode, expiryTime := helpers.GenerateVerificationCode()
	user := models.User{
		Name: name,
		Email: email,
		Password: pass,
		Verification: verificationCode,  
		VerificationExpiry: expiryTime,
	}

	maildata := models.MailData{
		To:       "admin@gmail.com",
		From:      user.Email,
		Subject:  "verification",
		Content:  fmt.Sprintf("Your verification code is %s",user.Verification),
		Template: "",
	}

	m.App.MailChan <-maildata
	err = m.DB.CreateUser(user)
	if err != nil {
		log.Println("error in handlers createuser", err)
	}

	modUser, err := m.DB.GetVerifiedUserByEmail(email) 
	if err != nil {
		log.Println(err)
	}
	if !modUser.Verified {
		http.Redirect(w, r,"/verify",http.StatusSeeOther)
		log.Println("moved to verification Page")
	}
	http.Redirect(w, r,"/",http.StatusSeeOther)
	
}
	


func (m *Repository) Verify(w http.ResponseWriter, r *http.Request){
		// state := "hope is all"
		// a:= []byte(state)
		// w.Write(a)
		
		
}


func (m *Repository) PostVerify(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		email:= r.Form.Get("email")
		log.Println("v email from form is..",email)
		veriCode:= r.Form.Get("verificationCode")
		form:=forms.New(r.PostForm)
		//form.Required("email")
		//form.IsEmail(email)
		if !form.Valid(){
			resp := JsonResponse{
				OK:      false,
				Message: "email is missing or not valid",
			}
			out, err := json.MarshalIndent(resp, "", "   ")
			if err != nil {
				log.Println(err)
			}
			//w.Header().Set("Content-Type", "application/json")
			w.Write(out)
			
			return
		}
		modUser, err := m.DB.GetVerifiedUserByEmail(email)
		if err != nil {
			log.Println(err)
		}
		timeout:=modUser.VerificationExpiry
		vCode := modUser.Verification
		log.Println(vCode ," == ",veriCode)
			log.Println(timeout)
			log.Println("time of sending v-code",timeout)
			log.Println("time now",time.Now())
			log.Println("time of x",time.Now().After(timeout.Add(5*time.Minute)))
			log.Println("time of expiry",timeout.Add(5*time.Minute))
			
		log.Println("v codddeee from db is..",vCode,modUser.Email)
		if time.Now().After(timeout.Add(5*time.Minute))||veriCode!=vCode {
			
			resp := JsonResponse{
				OK:      false,
				Message: "Verification code is expired or invalid",
			}
			out, err := json.MarshalIndent(resp, "", "   ")
			if err != nil {
				log.Println(err)
			}
			//w.Header().Set("Content-Type", "application/json")
			w.Write(out)
			return
		}
		//TODO: Update database to mark user as verified
		resp := JsonResponse{
			OK:      false,
			Message: "user is verified",
		}
		out, err := json.MarshalIndent(resp, "", "   ")
		if err != nil {
			log.Println(err)
		}
		//w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		

}

func (m *Repository)HomePage(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"arcform.page.html",&models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository)PostHome(w http.ResponseWriter, r *http.Request){
	log.Println("entered the post handler")
	err:=r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	screenSize := r.Form.Get("Screen_size")
	storageSize := r.Form.Get("storage_size")
	gen := r.Form.Get("gen")
	ram := r.Form.Get("ram_size")
	cpu := r.Form.Get("cpu")
	resolution := r.Form.Get("resolution")
	form := forms.New(r.PostForm)
	form.Required("Screen_type","storage_size","gen")
	// if !form.Valid(){
	// 	log.Println("somee required fields are missing")
	// 	render.Template(w,r,"arcform.page.html",&models.TemplateData{
	// 		Form: forms.New(nil),
	// 	})
	// }

	
	
	hp := helpers.NewPC()
	hp.ScreenSize, _ = strconv.Atoi(screenSize)
	hp.Gen, _ = strconv.Atoi(gen)
	hp.StorageSize, _ = strconv.Atoi(storageSize)
	hp.RamSize, _ = strconv.Atoi(ram)
	hp.Resolution, _ = strconv.Atoi(resolution)
	hp.Cpu, _ = strconv.Atoi(cpu)
	
	price := hp.EstimatePrice()
	hp.Price = price
	log.Println("printing thissssss:",gen,storageSize,price,hp.Gen)
	m.App.Sessions.Put(r.Context(),"hp",hp)
	http.Redirect(w,r,"/recommendation",http.StatusSeeOther)
	log.Println("made to go to recommendations")
	
}

func (m *Repository)RecommendationPage(w http.ResponseWriter, r *http.Request){
	retrievedHp:=m.App.Sessions.Get(r.Context(),"hp").(helpers.PcProperties)
	data := make(map[string]interface{})
	data["hp"]= retrievedHp

	render.Template(w,r,"generals.page.html",&models.TemplateData {
		//Form: forms.New(nil),
		Data: data,
	})
}

