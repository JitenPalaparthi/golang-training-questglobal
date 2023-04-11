var colors=["red","blue","green","orange"]
function demo() {
    try{
    var element = document.getElementById("p1")
    element.innerText = "Hello World! I am from JavaScript"

   var tags= document.getElementsByTagName("p")
   let i=0
   for( var t of tags){
        t.innerText="I got changed"
        t.style.color=colors[i]
        i++
   }
   changeMe()
}catch(err){
alert(err)
    var e = document.getElementById("er")
   // e.style.visibility=visible;
    e.innerText=err
} 
}