import React from 'react';

function ReviewCard({ imageSrc, name, position, comment }) {
  return (
    <div className="relative flex w-full max-w-[26rem] flex-col rounded-xl bg-transparent bg-clip-border text-gray-700 shadow-lg">
      <div className="relative flex items-center gap-4 pt-0 pb-8 mx-0 mt-4 overflow-hidden text-gray-700 bg-transparent shadow-none rounded-xl bg-clip-border">
        <img
          src={imageSrc}
          alt={name}
          className="relative inline-block h-[58px] w-[58px] !rounded-full object-cover object-center"
        />
        <div className="flex w-full flex-col gap-0.5">
          <div className="flex items-center justify-between">
            <h5 className="block font-sans text-xl antialiased font-semibold leading-snug tracking-normal text-blue-gray-900">
              {name}
            </h5>
            <div className="flex items-center gap-0.5">
              {[...Array(5)].map((_, index) => (
                <svg
                  key={index}
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  className="w-5 h-5 text-yellow-700"
                >
                  <path
                    fillRule="evenodd"
                    d="M10.788 3.21c.448-1.077 1.976-1.077 2.424 0l2.082 5.007 5.404.433c1.164.093 1.636 1.545.749 2.305l-4.117 3.527 1.257 5.273c.271 1.136-.964 2.033-1.96 1.425L12 18.354 7.373 21.18c-.996.608-2.231-.29-1.96-1.425l1.257-5.273-4.117-3.527c-.887-.76-.415-2.212.749-2.305l5.404-.433 2.082-5.006z"
                    clipRule="evenodd"
                  ></path>
                </svg>
              ))}
            </div>
          </div>
          <p className="block font-sans text-base antialiased font-light leading-relaxed text-blue-gray-900">
            {position}
          </p>
        </div>
      </div>
      <div className="p-4 mb-6 text-justify">
        <p className="block font-sans text-base antialiased font-normal leading-relaxed text-inherit">
          "{comment}"
        </p>
      </div>
    </div>
  );
}

function Reviews() {
  const reviews = [
    {
      imageSrc: 'https://media.licdn.com/dms/image/D5603AQHQXM5XlUqjTg/profile-displayphoto-shrink_800_800/0/1705721021390?e=1718841600&v=beta&t=_IQb_Fg1Tf_BSN1QMhRWIiurjiiGSEWoCwzlf1QSPHk',
      name: 'Manu M Benakal',
      position: 'Co-founder of CosmicTech.io',
      comment: 'The website\'s sleek UI design makes navigation effortless, while it\'s vision for fostering community engagement shines through every aspect.',
    },
    
    {
      imageSrc: 'https://media.licdn.com/dms/image/D5603AQFfHKWK90jPyQ/profile-displayphoto-shrink_800_800/0/1692205618654?e=1718841600&v=beta&t=DFOc3PLUENQ8p03li7vRN3wVd11t_p0y019Icdd46Qo',
      name: 'Mohit R',
      position: 'Associate Software Engineer @Betsol',
      comment: 'Impressed by the intuitive user interface, which seamlessly aligns with the project\'s vision of promoting inclusivity and accessibility.',
    },

    {
      imageSrc: 'https://avatars.githubusercontent.com/u/64711780?v=4',
      name: 'Varun S Athreya',
      position: 'Associate Software Engineer @Contentstack',
      comment: 'The website\'s visually appealing layout not only captivates but also effectively communicates the project\'s noble vision of empowering individuals through technology.',
    },
  ];

  return (
    <div id="reviews" className=" ">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h2 className="text-3xl font-extrabold text-gray-900 sm:text-4xl p-auto mb-10 pt-8">
            Reviews
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-10">
            {reviews.map((review, index) => (
              <ReviewCard key={index} {...review} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Reviews;
