import React from 'react';
import './Navbar.css'

const Navbar = () => {
    return (
        <div className='Nav'>
            <nav>
                <h1>Logo</h1>

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
