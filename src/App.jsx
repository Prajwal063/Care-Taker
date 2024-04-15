import React, { useState, useEffect } from "react";
import { Suspense } from "react";
import { PropagateLoader } from "react-spinners";
import Home from "./pages/Home";

const App = () => {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => {
      setLoading(false);
    }, 1000);
    return () => clearTimeout(timer);
  }, []);

  return (
    <Suspense
      fallback={
        <div className="h-screen w-full flex items-center justify-center">
          <PropagateLoader color="#1679AB" height={50} width={6} radius={4} margin={6} />
        </div>
      }
    >
      {loading ? (
        <div className="h-screen w-full flex items-center justify-center">
          <PropagateLoader color="#1679AB" height={50} width={6} radius={4} margin={6} />
        </div>
      ) : (
        <Home />
      )}
    </Suspense>
  );
};

export default App;
