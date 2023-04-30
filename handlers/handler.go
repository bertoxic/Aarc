package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	screenType := r.Form.Get("Screen_type")
	storageSize := r.Form.Get("storage_size")
	ramSize := r.Form.Get("gen")

	log.Println("printing thissssss:",ramSize,storageSize,screenType)
	

	render.Template(w,r,"arcform.page.html",&models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository)General(w http.ResponseWriter, r *http.Request){
	err:=r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	render.Template(w,r,"generals.page.html",&models.TemplateData{
		Form: forms.New(nil),
	})
}

