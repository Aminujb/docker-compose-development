import "./App.css";

import Reports from "./components/ReportTable";
import CreateReport from "./components/ReportForm";

import { Switch, Route, Link } from "react-router-dom";

const NoMatch = () => <h1>The page you requested was not found.</h1>;
const routes = (
  <div>
    <Switch>
      <Route path="/create-report" component={CreateReport} exact />
      <Route path="/" component={Reports} exact />
      <Route path="*" component={NoMatch} />
    </Switch>
  </div>
);

const Nav = () => (
  <>
    <div style={{ marginBottom: "10px"}}>
      <Link to="/">View Report</Link>
      &nbsp;&nbsp;
      <Link to="/create-report">Create Report</Link>
    </div>
  </>
);

const App = () => (
  <div style={{ margin: "10px" }}>
    <Nav/>
    {routes}
  </div>
);

export default App;
