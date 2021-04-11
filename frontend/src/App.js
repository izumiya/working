import './App.css';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';
import {Navbar, Nav, Form, Button, FormControl} from "react-bootstrap";
import CoffeeList from './CoffeeList.js';
import Admin from './Admin.js';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
    return (
        <Router>
            <div className="App">
                <Navbar bg="light" expand="lg">
                    <Navbar.Brand href="#home">Coffee Shop</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav"/>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="mr-auto">
                            <Nav.Link href="/">Home</Nav.Link>
                            <Nav.Link href="/admin">Admin</Nav.Link>
                        </Nav>
                        <Form inline>
                            <FormControl type="text" placeholder="Search" className="mr-sm-2"/>
                            <Button variant="outline-success">Search</Button>
                        </Form>
                    </Navbar.Collapse>
                </Navbar>
                <Switch>
                    <Route path="/admin"><Admin/></Route>
                    <Route path="/"><CoffeeList/></Route>
                </Switch>
            </div>
        </Router>
    );
}

export default App;
