import React from "react";
import "./Landing.css";
import landingpic from "./landing.png";

function Landing(props) {
  return (
    <div>
      <div class="maintext">
        <h1>Make School</h1>
        <h1>Courses</h1>
        <h1>Told By Students</h1>
      </div>

      <div class="pls">
        <button class="button button4">VIEW COURSES</button>
      </div>
      <div class="landing">
        <img src={landingpic} alt="" />
      </div>
      <img
        src="https://res.cloudinary.com/erica-naglik/image/upload/v1564602895/Screen_Shot_2019-06-07_at_7.46.14_PM_cbfogy.png"
        alt=""
        className="ms"
      />
    </div>
  );
}

export default Landing;
