import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { Home } from "./pages/home";

import { Login } from "./pages/login";

import { Error } from "./pages/404";
import { Footer } from "./components/Communs/Footer";
import { NavBar } from "./containers/Communs/NavBar";
import { Collection } from "./pages/collection";
import { NavSpace } from "./components/Communs/NavBar";

export const Index = () => {
	return (
		<Router>
			{true && (
				<>
					<NavBar />
					<NavSpace />
				</>
			)}
			<Switch>
				<Route path="/" exact component={Home} />
				<Route path="/login" exact component={Login} />
				<Route path="/collection" exact component={Collection} />
				<Route>
					<Error value={404} />
				</Route>
			</Switch>
			{true && <Footer />}
		</Router>
	);
};
