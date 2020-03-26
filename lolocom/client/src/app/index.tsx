import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { Home } from "./pages/home";

import { Login } from "./pages/login";

import { Error } from "./pages/404";
import { Footer } from "./containers/Communs/Footer";
import { NavBar } from "./containers/Communs/NavBar";

export const Index = () => {
	return (
		<Router>
			{true && <NavBar />}
			<Switch>
				<Route path="/" exact component={Home} />
				<Route path="/login" exact component={Login} />
				<Route>
					<Error value={404} />
				</Route>
			</Switch>
			{true && <Footer />}
		</Router>
	);
};
