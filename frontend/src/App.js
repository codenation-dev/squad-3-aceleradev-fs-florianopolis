import React,{Fragment} from 'react';
import Nav from './Components/Nav';
import Login from './Components/Login';
import Dashboard from './Components/Dashboard';
import Register from './Components/Register';
import {Switch, Route , Redirect} from 'react-router-dom';
import {connect} from 'react-redux';
import UnathRedirect from './Helpers/UnauthorizedRedirect'
import './static/app.css';

const LogicalGate = (props) => (
  props.isLogged.value?<Redirect to="/dashboard"/>:<Redirect to="/login" />
  )

const DashboardRoute = () => (<UnathRedirect Auth={<Dashboard />} Unath={<Redirect to="/login"/>} />)
const RegisterRoute = () => (<UnathRedirect Auth={<Register />} Unath={<Redirect to= "/login" />} />)
const WarnRoute = () => (<UnathRedirect Auth={<Dashboard />} Unath={<Redirect to="/login" />} />)
const App = (props) => {
  return (
    <Fragment>
      {console.log(props.isLogged)}
    <div>
    <Route path="/" component={Nav}/>
    </div>
    <main>
    <Switch>
    <Route path="/" exact component={() => <LogicalGate isLogged={props.isLogged} />} />
    <Route path="/login" exact component={Login}/>
    <Route path="/dashboard" exact component={DashboardRoute} />
    <Route path="/mailregister" exact component={RegisterRoute}/>
    <Route path="/warn" exact component={WarnRoute} />
    </Switch>
    </main>
    </Fragment>);
}

const mapStateToProps = state => {
  return {
    isLogged: state
  }
}

export default connect(mapStateToProps,null)(App);
