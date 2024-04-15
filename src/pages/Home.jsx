import React from "react";

import Header from "../components/header/Header";
import Footer from "../components/footer/Footer";
import Hero from "../components/hero/Hero";
import About from "../components/about/About";
import scroll from "../assets/scroll.png"
import What from "../components/what/What";
import Events from "../components/events/Events";
import Reviews from "../components/reviews/Reviews";
import Contact from "../components/contact/Contact";

const Home = () => {

  function scrollToTop() {
    window.scrollTo({top: 0, behavior: "smooth"})
  }

  return (
    <>
      <Header/>
      <Hero/>
      <About/>
      <What/>
      <Events/>
      <Reviews/>
      <Contact/>
      <Footer/>

      
      <div className="fixed bottom-4 right-4">
      <button 
          onClick={scrollToTop}
          className="px-0 py-2 rounded-full">
          <img
              className="w-[40px] h-[40px] md:w-[40px] md:h-[40px]"
              src={scroll}
              alt="Scroll icon"/>
        </button>
      </div>
    </>
  );
};

export default Home;
