import React, { useState } from 'react';
import logo from "../../assets/care taker logo.png"


const navItems = [
  {
    id: 0,
    name: "About",
  },
  {
    id: 1,
    name: "Services",
  },
  {
    id: 2,
    name: "Events",
  },
  {
    id: 3,
    name: "Donate",
  },
  {
    id: 4,
    name: "Reviews",
  },
  {
    id: 5,
    name: "Contact",
  },
  {
    id: 6,
    name: "Volunteer",
  },
  
];

const Header = () => {

  const renderNavItems = (navItems) => {
    const scrollToSection = (id) => {
      const section = document.getElementById(id);
      if (section) {
        section.scrollIntoView({ behavior: "smooth" });
      }
    };
  
    return navItems.map(item => (
      <li key={item.id}>
        <a 
          onClick={() => {
            toggleMenu(item.name);
            scrollToSection(item.name.toLowerCase());
          }}
          href={`#${item.name}`} 
          className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0  "
        >
          {item.name}
        </a>
      </li>
    ));
  };  

  const [isOpen, setIsOpen] = useState(false);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  return (
    <nav className="bg-white fixed w-full z-20 top-0 start-0 border-b border-gray-200">
      <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
        <a href="/" className="flex items-center space-x-3 rtl:space-x-reverse">
          <img src={logo} className="h-10" alt="Care-Taker Logo" />
        
        </a>

        <div className="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
          <button type="button" class="bg-blue-500 hover:bg-blue-600 text-white focus:ring-4 focus:outline-none focus:ring-blue-300  rounded-full text-sm px-4 py-2 text-center">Get Involved</button>

          <button data-collapse-toggle="navbar-sticky" 
          type="button" 
          className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200" 
          aria-controls="navbar-sticky" 
          aria-expanded="false" 
          onClick={toggleMenu}>
            <span className="sr-only">Open main menu</span>
            <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
              <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 1h15M1 7h15M1 13h15"/>
            </svg>
          </button>
        </div>
        <div className={`items-center justify-between w-full md:flex md:w-auto md:order-1 ${isOpen ? 'block' : 'hidden'}`} id="navbar-sticky">
          <ul className="flex flex-col p-4 md:p-0 mt-4 font-medium border rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white">
            {renderNavItems(navItems)}
          </ul>
        </div>
      </div>
    </nav>

  );
};

export default Header;

