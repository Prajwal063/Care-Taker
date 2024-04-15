import React, { useState } from 'react';

function renderCard(imageSrc, title, description) {
  return (
    <div className="relative flex flex-col text-gray-700 bg-white shadow-md bg-clip-border rounded-xl w-96 mb-8 sm:mb-8 pb-6 sm:pb-8">
      <div className="relative h-56 mx-4 -mt-6 overflow-hidden text-white shadow-lg bg-clip-border rounded-xl bg-blue-gray-500 shadow-blue-gray-500/40">
        <img src={imageSrc} alt="card"/>
      </div>
      <div className="p-4 pb-0">
        <h5 className="block mb-2 font-sans text-xl text-center antialiased font-semibold leading-snug tracking-normal text-blue-gray-900">{title}</h5>
        <p className="block font-sans text-base antialiased font-light leading-relaxed text-inherit">{description}</p>
      </div>
    </div>
  );
}

function What() {
  const [currentPage, setCurrentPage] = useState(1);

  // Sample data for cards
  const cardsData = [
    {
      imageSrc: "https://images.unsplash.com/photo-1540553016722-983e48a2cd10?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=800&amp;q=80",
      title: "Visit Care Homes",
      description: "Discover and engage with our network of care homes and orphanages, offering support and fostering community for their residents."
    },
    {
      imageSrc: "https://images.unsplash.com/photo-1540553016722-983e48a2cd10?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=800&amp;q=80",
      title: "Celebrate Events",
      description: "We can organize events such as fundraisers, birthday, awareness campaigns, or recreational activities for the residents."
    },
    {
      imageSrc: "https://images.unsplash.com/photo-1540553016722-983e48a2cd10?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=800&amp;q=80",
      title: "Volunteer Sign-Up",
      description: "You can sign up as volunteers and offer their time and skills to help out at orphanages and old age homes."
    },
    {
      imageSrc: "https://images.unsplash.com/photo-1540553016722-983e48a2cd10?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=800&amp;q=80",
      title: "Resource Sharing",
      description: "Share informative articles, guides, and resources on topics related to caregiving, mental health, and well-being for both caregivers and residents."
    },
    {
      imageSrc: "https://images.unsplash.com/photo-1540553016722-983e48a2cd10?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=800&amp;q=80",
      title: "Virtual Tours",
      description: "Provide virtual tours of orphanages and old age homes, giving donors and volunteers insight into their facilities and operations."
    },
    {
      imageSrc: "https://images.unsplash.com/photo-1540553016722-983e48a2cd10?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=800&amp;q=80",
      title: "Educational Resources",
      description: "Provide educational resources and programs for residents, such as literacy classes, vocational training, and skill development workshops."
    },
    
  ];

  // Calculate total number of pages based on number of cards and cards per page
  const totalPages = Math.ceil(cardsData.length / 3);

  // Calculate starting and ending index for the current page
  const startIndex = (currentPage - 1) * 3;
  const endIndex = Math.min(startIndex + 3, cardsData.length);

  // Render cards for the current page
  const renderedCards = cardsData.slice(startIndex, endIndex).map((card, index) => (
    <div key={index}>
      {renderCard(card.imageSrc, card.title, card.description)}
    </div>
  ));

  // Handle pagination click
  const handlePageClick = (pageNumber) => {
    setCurrentPage(pageNumber);
  };

  // Generate pagination buttons
  const paginationButtons = [];
  for (let i = 1; i <= totalPages; i++) {
    paginationButtons.push(
      <button
        key={i}
        onClick={() => handlePageClick(i)}
        className={`px-4 py-2 mx-1 rounded-md ${currentPage === i ? 'bg-gray-600 text-white' : 'bg-gray-200 text-gray-600 hover:bg-gray-300'}`}
      >
        {i}
      </button>
    );
  }

  return (
    <div id="services" className="pt-2 md:py-2 sm:py-2">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h2 className="text-3xl font-extrabold text-gray-900 sm:text-4xl p-auto mb-8 pt-8">
            Services
          </h2>
         
          <div className=" py-8 md-2 flex flex-col items-center justify-center">
            <div className="flex flex-wrap justify-center text-justify gap-8 md:gap-8 sm:pb-0">
              {renderedCards}
            </div>
            <div className="flex mt-4">
              {paginationButtons}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default What;
