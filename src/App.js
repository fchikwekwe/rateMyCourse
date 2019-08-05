import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom'

import SignUp from './SignUp'
import LogIn from './LogIn'
import NavBar from './NavBar'
import Landing from './Landing'

function App() {
  return (
    <div className="App">
    <BrowserRouter>
      
        <NavBar/>
        
        <Route path="/" exact component={Landing} />
        <Route path="/login" exact component={LogIn} />
        <Route path="/signup" exact component={SignUp} />
        
        </BrowserRouter>
    </div>
  );
}

export default App;
