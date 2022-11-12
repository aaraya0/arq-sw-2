
import React, {useEffect, useState}  from "react";

import Cookies from 'universal-cookie';
import "./Results.css";
const cookies = new Cookies();
async function getSearch(query){
    return (fetch('http://localhost:8983/solr/publicaciones/select?&defType=lucene&indent=true&q=description:"'+query+'"%0Atitle:"'+query+'"&q.op=OR').then(response => response.json()));
     }

async function getStuff(){
    let items= await getSearch(cookies.get("busqueda_limpia"))
    
}
   
function Results(){
   

}
export default Results;