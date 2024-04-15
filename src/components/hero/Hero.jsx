import React from "react";
import image1 from "../../assets/try.png";
import 'react-responsive-carousel/lib/styles/carousel.min.css';

const Hero = () => {
  return (
    <div id="hero" className="bg-sky-100 flex flex-col md:flex-row items-center justify-center py-8 md:py-12">
      {/* Left Side */}
      <div className="md:w-1/2 flex flex-col items-start justify-center px-10 pt-12 md:pt-0">
        <h1 className="text-3xl font-bold text-gray-800 md:text-5xl mb-4 py-6">
          A nurturing, supportive community platform.
        </h1>
        <p className="text-lg text-gray-700 md:text-xl mb-10 text-justify">
          Care Taker connects people to orphanages and old age homes, fostering positive impact through visits, donations, and events.
        </p>
        <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-7 rounded-full transition duration-300 ease-in-out md:text-lg">
          Get Involved
        </button>
      </div>

      {/* Right Side */}
      <div className="md:w-1/2 flex flex-col justify-center items-center pt-10">
        <div className="rounded-lg overflow-hidden mb-6">
          <img src={image1} alt="Orphanage" className="w-full h-auto" />
        </div>
        <div className="text-gray-700 text-xs md:text-base sm:text-xs">
          <a href="https://prajwalp06.netlify.app/" className="hover:underline">Prajwal P</a> | Full Stack Developer | Bengaluru
        </div>
      </div>
    </div>
  );
};

export default Hero;
