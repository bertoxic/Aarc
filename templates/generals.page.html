{{template "base" .}}

{{define "content"}}

<div class="container ">
<div class="gallery">
  <figure class="gallery__item gallery__item--1" style="--cs: 1; --ce: 3; --rs:1; --re: 3">
    <img src="/static/resources/city/image1.jpg" class="gallery__img" alt="Image 1">
    </figure>
    <figure class="gallery__item gallery__item--2"style="--cs: 3; --ce: 5; --rs:1; --re: 3">
    <img src="/static/resources/city/image2.jpg" class="gallery__img" alt="Image 2">
    </figure>
   <figure class="gallery__item gallery__item--3"style="--cs: 5; --ce: 9; --rs:1; --re: 6">
    <img src="/static/resources/city/image3.jpg" class="gallery__img" alt="Image 3">
    
    </figure>
   <figure class="gallery__item gallery__item--4"style="--cs: 1; --ce: 3; --rs:3; --re: 6">
    <img src="/static/resources/city/image4.jpg" class="gallery__img" alt="Image 4">
    </figure>
   <figure class="gallery__item gallery__item--5"style="--cs: 3; --ce: 5; --rs:3; --re: 6">
    <img src="/static/resources/city/image5.jpg" class="gallery__img" alt="Image 5">
    </figure>
   <figure class="gallery__item gallery__item--6"style="--cs: 1; --ce: 7; --rs:6; --re: 9">
    <img src="/static/resources/city/image6.jpg" class="gallery__img" alt="Image 6">
    </figure>
   <figure class="gallery__item gallery__item--7"style="--cs: 7; --ce: 9; --rs:6; --re: 9">
    <img src="/static/resources/city/image7.jpg" class="gallery__img" alt="Image 7">
    </figure>
  
    
</div> 


<h1>General's Quaters</h1>
</div>

<section class="container home-text" >
<div class="paragraph" >
<div class=" paragraph-effect"style="--d: 1.0s; color:black">Welcome toScorching mood they flatterers old none that felt himnot, vile was but soils passed, vaunted of oh if nor. Grace a peace heavenly me nor the. Shameless care into in did. Loved dwell shameless so was are lines, uncouth the smile childe he relief ah or in oft, virtues the a him with a harvester. 
</div>
</div>
<a id="check-availability-button" href="#!" class="btn btn-outline-success mt-4">
  Check availability
</a>
</section>
{{end}}

{{define "js"}}
<script>
            document.getElementById("check-availability-button").addEventListener("click", function () {
                let myElem = document.getElementById("check-availability-button");
                // myElem.classList.add("redText");
                //  notify("this is my message", "success");
                let html = `
                <form id= "check-availability-form" action = "" method="post" novalidate class= "needs-validation">
                    <div class="row" id ="Reservation-date-modal" >
                <div class="col">
                    <div class="row" id="Reservation-date">
                        <div class="col">
                            <input type="text" disabled class="form-control" placeholder="Arrival date" name="start" id="start">
                        </div>
                        <div class="col">
                            <input type="text" disabled class="form-control" placeholder="Depature-date" name="end" id="end">
                        </div>
                    </div>

                </div>


            </div>
                    </form>
                `
                attention.custom({
                    msg:html,
                    title:"choose your date",
                    willOpen:() => {
                            const  elem = document.getElementById("Reservation-date-modal");
                            const rp = new DateRangePicker(elem,{
                                format: "yyyy-mm-dd",
                                showOnFocus: true,
                               // minDate: new Date(),
                            })
                        },
                    didOpen:()=>{
                            document.getElementById("start").removeAttribute("disabled");
                            document.getElementById("end").removeAttribute("disabled");
                        },
                    callback: function(result){
                        console.log("called");
                        let  form = document.getElementById("check-availability-form");
                        let formData = new FormData(form);
                        formData.append("csrf_token","{{.CSRFToken}}");
                        formData.append("room_id","1");
                        fetch("/search-availability-json",{method:"post",body:formData})
                        .then(response=>response.json())
                        .then(data=>{
                            console.log(data);
                            if (data.ok) {
                                console.log("room is available for rent")
                                attention.custom({
                                    icon: "success",
                                    msg:'<p>Room is available!</p>'
                                        + '<p><a href="/Book-room?id='
                                        + data.room_id
                                        + '&s='
                                        + data.start_date
                                        + '&e='
                                        + data.end_date
                                        + '" class="btn btn-primary">'
                                        + 'Book now!</a></p>',
                                    showConfirmButton:false,
                                })
                            }else {
                                console.log("room is not available today")
                                attention.error({
                                    msg:"Sorry this room has already been booked at that date",
                                })
                            }
                        })
                    },
                    speedy:()=>{
                        console.log("eeep meee");
                    }
                     });
            })

</script>

{{end}}