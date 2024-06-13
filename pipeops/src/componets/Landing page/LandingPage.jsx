import React from 'react';
import './LandingPage.css';
import Navbar from '../Navbar/Navbar';

const LandingPage = () => {
    return (
        <>
            <Navbar />
            <section>
                <div style={{
                    width:'100%',
                    height:'90vw',
                    backgroundSize:'cover',
                    backgroundPosition:'center',
                    
                    backgroundImage: 'url(./Assets/pix/james-donaldson-toPRrcyAIUY-unsplash (1).jpg)'
                    }} 
                    className="landing">
                    hello
                </div>
            </section>
        </>
    );
}

export default LandingPage;
