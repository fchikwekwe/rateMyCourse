import React, { Component } from "react";
import "./LogIn.css";
import { Link } from "react-router-dom";
const Input = props => (
  <input
    className="c-input"
    type={props.type}
    placeholder={props.placeholder}
    onKeyUp={props.onKeyUp}
  />
);

const Button = props => (
  <button className="c-button" onClick={props.onClick}>
    {props.text}
  </button>
);

class LogIn extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: { value: "" },
      password: { value: "" }
    };
  }

  handleInput = (param, e) => {
    let value = e.target.value;

    console.log(param);

    if (param === "name") {
    }
  };

  render() {
    return (
      <div>
        <section id="signup">
          <div className="container ">
            <div className="row">
              <div className="col-md-10 mx-auto">
                <div className="card">
                  <div className="row mr-0 ml-0 d-flex h-100">
                    <img
                      src="https://res.cloudinary.com/erica-naglik/image/upload/v1564606802/skyline_vqpy6p.png"
                      className="img-fluid"
                      width="860px"
                    />

                    <div className="col-md-6 form">
                      <div className="card-title">
                        <img
                          src="https://res.cloudinary.com/erica-naglik/image/upload/v1564602895/Screen_Shot_2019-06-07_at_7.46.14_PM_cbfogy.png"
                          className="img-ms logo"
                        />
                        <h2>Log In</h2>
                        <p>
                          Welcome Back! Hope you're having a great day at Make
                          School!
                        </p>
                      </div>
                      <div className="card-body">
                        <form className="form">
                          <Input
                            type="text"
                            placeholder="Make School Email"
                            value={this.state.email.value}
                            onKeyUp={e => this.handleInput("email", e)}
                          />

                          <Input
                            type="password"
                            placeholder="Create password"
                            value={this.state.password.value}
                            onKeyUp={e => this.handleInput("password", e)}
                          />
                          <div class="plsgod">
                            <Button text="Log In" onClick={this.handleSubmit} />
                          </div>
                          <div class="plsbutton">
                            <Link to="/signup" className="navbar__link">
                              Need to sign up? Click here
                            </Link>
                          </div>
                        </form>
                      </div>
                    </div>
                  </div>
                </div>{" "}
                {/* End of card */}
              </div>
            </div>
          </div>
        </section>
      </div>
    );
  }
}

export default LogIn;
