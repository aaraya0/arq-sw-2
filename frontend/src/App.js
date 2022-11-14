import React from 'react';
import './App.css';
import Search from './components/Search';
import Results from './components/Results';
import NavBar from './components/NavBar';

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
function App() {
  return (
    <>
    <Router>
    
    <NavBar/>
    <Routes>
    <Route exact path="/" element={<Search/>}/>
    <Route exact path="/results" element={<Results/>}/>
    </Routes>
 
    
    
    </Router>
      
  
    </>
  );
}

export default App;