import React, {useState} from 'react'
import {Link} from 'react-router-dom'
import './NavBar.css';
import image from  "../tree-and-roots.png"

import Cookies from 'universal-cookie';



const cookies = new Cookies();

function NavBar() {
const [click, setClick]=useState(false);
const [setButton]=useState(true);
const handleClick=()=>setClick(!click);
const closeMobileMenu=()=>setClick(false);
const showButton =()=> {
if(window.innerWidth <=960){
setButton(false);
} else {
setButton(true);
}
};
window.addEventListener('resize', showButton);
var username=  cookies.get("name");

window.addEventListener('resize', showButton);
var username=  cookies.get("name");

        return (
        <>
        <nav className="NavBar" >
        <div className="navbar-container">
        <Link to="/" className="navbar-logo">
        <img src={image} className="imagen"/>
        </Link>
        <div className='menu-icon' onClick={handleClick}>
        <i className={click ? 'fas fa-times': 'fas fa-bars'}/>
        </div>
        <ul id="lista" className={click ? 'nav-menu active' : 'nav-menu'}> 
        <li className='nav-item'>
        <Link to='/publish' className='nav-links' onClick={closeMobileMenu}>
        Publicar
        </Link>
        </li>
        <li className='nav-item'>
        <Link to='/login' className='nav-links' onClick={closeMobileMenu}>
        Iniciar Sesión
        </Link>
        </li>
       
        </ul>

        </div>
        </nav>
        </>
        )
}

export default NavBar