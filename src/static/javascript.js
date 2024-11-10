//const element = document.querySelector('.commit.first.commit0');
var element = document.getElementsByClassName('first')[0];
// do{
//     element = document.getElementsByClassName('first');
// } while(element.length == 0)
console.log(element);
//const element = document.getElementById("this")
//const element = document.querySelector('.commit.first.commit0'); // Use querySelector with a class selector
if (element){
    alert(element)
}
element.addEventListener('mouseover', () => {
element.style.fill = "red"; // Change background color on hover   
})