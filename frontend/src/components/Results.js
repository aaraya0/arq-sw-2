
import React, {useEffect, useState}  from "react";

import Cookies from 'universal-cookie';
import "./Results.css";
const cookies = new Cookies();
async function getSearch(query){
    return (fetch("http://localhost:8983/solr/publicaciones/select?" + query + "=&defType=lucene&omitHeader=true&indent=true&q.op=AND&q=*%3A*")
    .then(response => response.json()));
     }

async function getStuff(){
    let items= await getSearch(cookies.get("busqueda_limpia"))
    console.log(items)
           
            return items
}
   


function gopath(path){
    window.location = window.location.origin + path
  }

  

class Results extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
        items: [],
        DataisLoaded: true
        };
        }
       
        

        componentDidMount() {
            this.setState(getStuff())
            }

            render() {
            const { DataisLoaded, items} = this.state;
                if (!DataisLoaded) return <div>
                <h1> Please wait... </h1> </div> ;
            
            
            const publi= items.map((item) => (
            
                <div id = { item.id } className="item">
                    <div id="titulo">{ item.title}</div>
                    <div id="desc"> { item.description }</div>	
                    <div id="seller"> { item.seller }</div>	
                    <div id="city">{ item.city }</div>
                    <div id="state">{ item.state }</div>
              </div>
               
            
            ))
         
            return (
                <div >
                
               
               
                <div className="public">{publi}</div>
                </div>
                )}
            
            

}
export default Results;