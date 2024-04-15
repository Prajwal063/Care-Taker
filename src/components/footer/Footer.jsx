import React from "react";
import logo from "../../assets/care taker logo.png"

const Footer = () => {
  return (
    <footer className="bg-white rounded-lg shadow dark:bg-gray-900 m-4">
      <div className="w-full max-w-screen-xl mx-auto p-4 md:py-8">
          <div className="sm:flex sm:items-center sm:justify-between">
              <a href="/" className="flex items-center justify-center mb-4 sm:mb-0 space-x-3 rtl:space-x-reverse ">
                  <img src={logo} className="h-10" alt="Care Taker Logo" />
              </a>
              <ul className="flex items-center justify-center text-sm text-gray-500 sm:mb-0 dark:text-gray-400">
                <span className="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Created with ❤️
                by <a className="hover:underline md:me-2" href = "https://prajwalp06.netlify.app/"> Prajwal P</a></span>   
              </ul>
          </div>
          <hr className="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
          <span className="block text-gray-500 sm:text-center justify-center dark:text-gray-400">© 2024 <a href="/" className="hover:underline">Care Taker™</a>. All Rights Reserved.</span>
      </div>
    </footer>
  );
};

export default Footer;
