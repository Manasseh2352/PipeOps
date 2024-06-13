import React from 'react';
import './Navbar.css'

const Navbar = () => {
    return (
        <div className='Nav'>
            <nav>
                <img src="./Assets/pix/submark-black.png" alt="" />

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
