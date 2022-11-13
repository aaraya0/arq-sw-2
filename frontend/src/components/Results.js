import React, { useState , useEffect} from "react";

import Cookies from "universal-cookie";



const cookies = new Cookies();

async function getItems(){
    let query= cookies.get("busqueda_limpia")
  return await fetch('http://localhost:8983/solr/publicaciones/select?&defType=lucene&indent=true&q=description:"'+query+'"title:"'+query+'"&q.op=OR')
.then(response => response.json())
}
//fetch("htttp://localhost:8090/search"+query)





function showItems(items){
  return items.map((item) =>

   <div obj={item} key={item.id} className="item">
   
    <a className="title">{item.title}</a>
    <a className="price"> {"$" + item.price}</a>

  
    <div>
      <a className="location">{item.city},</a>
    </div>
    <div>
      <a className="description">{item.description}</a>
    </div>
    <div className="sellerBlock">
      <a className="seller">{item.seller}</a>
    </div>

   </div>
 )
}



function Results() {
  const [items, setItems] = useState([])
 
  getItems().then(response => setItems(response))

  
    
  

  return (
    <div className="home">
  

     

      <div id="main">
       {showItems(items)}

      </div>
    </div>
  );
}

export default Results;