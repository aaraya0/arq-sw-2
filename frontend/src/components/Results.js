
import React, {useEffect, useState}  from "react";

import Cookies from 'universal-cookie';
import "./Results.css";
const cookies = new Cookies();
async function getSearch(query){
    return (fetch('http://localhost:8983/solr/netflix_core/select?indent=true&json={"query":{"dismax":{"df":"title"%2C"query":"'+query+'"}}}&q.op=OR&q=*:*', {method:"GET",
     mode: 'no-cors'}).then(response => response.json()));
     }

async function getStuff(){
    let items= await getSearch(cookies.get("busqueda_limpia"))
  
    let chain=''
    let a=items.split('"docs":[')
    let b=a[1]
    let c=b.split("}")
    for (let i = 0; i < c.length; i++){
       let d= c.split('"title:"[')
       let e = d[1]
       let f= e.split("]")
       let g = f[0]
       chain=`${chain}, ${g}`
    }
cookies.set("ayudaa", chain)

}
   
function Results(){
    getStuff()

}
export default Results;