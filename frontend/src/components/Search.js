import React from "react";
import "./Search.css";
import Cookies from 'universal-cookie';
import {useState} from 'react';
const cookies = new Cookies();

function gopath(path){
    window.location = window.location.origin + path
  }
  async function getSearch(query){
   return (fetch('http://localhost:8983/solr/netflix_core/select?indent=true&json={"query":{"dismax":{"df":"title"%2C"query":"'+query+'"}}}&q.op=OR&q=*:*', {method:"GET",
    mode: 'no-cors'
}).then(response => response.json()))
    }
  
async function SearchByQuery(message){

    let result= await getSearch(message)
    
      cookies.set ("ids",`${result.title},` ,'/' )

}
function Search(){

   
  
   //var {query} = document.forms[0];
const renderForm = (
    <div id="cover">
  <form method="get" action="">
    <div class="tb">
      <div class="td">
        <input type="text" id="search_query" placeholder="Buscar" required/></div>
      <div class="td" id="s-cover">
      <button  id="search_button" onClick={SearchByQuery(document.forms[0].value)} type="input">
          <div id="s-circle"></div>
          <span></span>
        </button>
      </div>
    </div>
  </form>
</div>
);

      return (
      <div className="app">
      <div className="search-form">

      {renderForm}

      </div>
      </div>
      );
    
}export default Search;