
import React, { useState , useEffect} from "react";

import Cookies from "universal-cookie";




const cookies = new Cookies();

async function getItems(){
 let query=cookies.get("busqueda_limpia")
  
  return await fetch("http://localhost:8090/search/"+query).then(response => response.json())
}


function goto(path){
  window.location = window.location.origin + path
}

function retry() {
  goto("/")
}

function parseField(field){
  if (field !== undefined){
    return field
  }
  return "Not available"
}



function showItems(items){
  return items.map((item) =>

   <div obj={item} key={item.id} className="item">
    
    <a className="title">{parseField(item.title)}</a>
    
    <div>
      <a className="location">{parseField(item.city)}</a>
     
    </div>
    <div>
      <a className="description">{parseField(item.description)}</a>
    </div>
    
   </div>
 )
}




function Results() {
  const [items, setItems] = useState([])
  const [needItems, setNeedItems] = useState(true)
 


  if(!items.length && needItems){
    getItems().then(response => setItems(response))
    setNeedItems(false)
  }







  return (
    <div className="home">
      <div className="topnavHOME">
      
    {showItems(items)}
        </div>

      </div>


  );
}

export default Results;