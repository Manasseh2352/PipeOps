import React from 'react';
import './Navbar.css'

const Navbar = () => {
    return (
        <div className='Nav'>
            <nav>
                <img src="./Assets/pix/submark-yellow.png" alt="" />
                <h1 style={{marginTop: '1rem',}}>SKUDOOSH</h1>

                <ul>
                    <li>Sign in</li>
                    <li>Support</li>
                    <li>About</li>
                </ul>
            </nav>
        </div>
    );
}

export default Navbar;
