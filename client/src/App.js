import logo from './logo.svg';
import Navbar from 'react-bootstrap/Navbar'

import { ApolloProvider } from '@apollo/react-hooks';
import ApolloClient from 'apollo-boost'

import Nav from 'react-bootstrap/Nav'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import Matrix from './components/Matrix';
import Reports from './components/Reports';
import TollBooth from './components/TollBooth';
import Home from './components/Home';

function App() {
  return (
    <Router>
      <>
  <Navbar bg="light">
    <Navbar.Brand href="/">Testing Shivansh</Navbar.Brand>
    <Nav>
      <Nav.Item>
    <Nav.Link><Link to="/matrix">Generate Matrix</Link></Nav.Link>
  </Nav.Item>
  <Nav.Item>
    <Nav.Link><Link to="/tollbooth">TollBooth</Link></Nav.Link>
  </Nav.Item>
  <Nav.Item>
    <Nav.Link><Link to="/reports">Reports</Link></Nav.Link>
  </Nav.Item>
    </Nav>
  </Navbar>
  <br />

</>,
      <div>
        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/matrix">
            <Matrix />
          </Route>
          <Route path="/reports">
            <Reports />
          </Route>
          <Route path="/tollbooth">
            <TollBooth />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
