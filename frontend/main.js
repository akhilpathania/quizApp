const show =()=> {
    const shows = document.querySelector("#signup")
    shows.classList.toggle("show")
    document.querySelector(".layer").style.display="block"
}
const hide = () => {
    const hides = document.querySelector("#signup")
    hides.classList.remove("show")
    document.querySelector(".layer").style.display="none"
}
function submitform(){
    var xhr=new XMLHttpRequest();
    xhr.open("POST","/submit-form");
    xhr.setRequestHeader("Content-Type","application/json");
    xhr.onreadystatechange = function(){
        if(xhr.readyState === 4 && xhr.status === 200){
            console.log(xhr.responseText);
        }
    };
    var name = document.getElementById("name").value;
    var email = document.getElementById("email").value;
    var data = JSON.stringify({"name":name, "email":email});
    xhr.send(data);
}