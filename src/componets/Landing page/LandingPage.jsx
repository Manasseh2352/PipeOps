import React from 'react';
import './LandingPage.css';

const LandingPage = () => {
    return (
        <>
            
            <section className="landing">
                <div style={{
                    position:'absolute',
                    width:'100%',
                    height:'100%',
                    backgroundSize:'cover',
                    backgroundPosition:'center',
                    filter:' blur(2px)',
                    zIndex: '-2',
                    backgroundImage: 'url("./Assets/pix/james-donaldson-toPRrcyAIUY-unsplash.jpg")'
                    }} 
                    >
                    
                </div>
                <div className="landpagetxt">
                    <h1>Welcome To</h1>
                    <img className='Txtlogo' src="./Assets/pix/primary-yellow.png" alt="" />
                    <div className="smalltxt">
                        <h2>Your Journey Companion </h2>
                        <p>Sharig Your journy experiences <br /> just got better</p>
                    </div>
                    <button>Get Started</button>
                </div>
            </section>
                    
            <section>
                <div className="subland">
                    <h3>Enjoy Every Experienc With Skudoosh</h3>

                    <div className="tabs">
                        <p><span>Trip Creation:</span>User's can ezaily plan and create their trips, setting destinatios and 
                            customizing itineraries
                        </p>
                        <div className="tabimg">
                                <img src="./Assets/pix/jeshoots-com-mSESwdMZr-A-unsplash (1).jpg" alt="" />
                           </div>
                    </div>
                    <div className="tabs">
                            <div className="tabimg">
                                <img src="./Assets/pix/jeshoots-com-mSESwdMZr-A-unsplash (1).jpg" alt="" />
                           </div>
                        <p><span>Trip Creation:</span>User's can ezaily plan and create their trips, setting destinatios and 
                            customizing itineraries
                        </p>

                    </div>
                </div>
            </section>
        </>
    );
}

export default LandingPage;
