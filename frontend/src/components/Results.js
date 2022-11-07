
import React, {useEffect, useState}  from "react";

import Cookies from 'universal-cookie';
import "./Results.css";
const cookies = new Cookies();
async function getSearch(query){
    return (fetch('http://localhost:8983/solr/publicaciones/select?city:"'+query+'"=&defType=lucene&description:"'+query+'"=&id:"'+query+'"=&indent=true&q.op=OR&q=*%3A*&seller:"'+query+'"=&title:"'+query+'"=', {method:"GET",
     mode: 'no-cors'}).then(response => response.json()));
     }

async function getStuff(){
    let items= await getSearch(cookies.get("busqueda_limpia"))
    return(
    <div>
    {items.map((item)=>
    <div>{item.title}</div>
    )}</div>)
}
   
function Results(){
    return(<div>{getStuff()}</div>)

}
export default Results;