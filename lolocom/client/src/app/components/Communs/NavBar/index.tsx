import React from "react";
import "./index.css";
import { LogIcon } from "../../../containers/Login";
import { SearchIcon } from "../../../containers/Communs/Search";
import { CartIcon } from "../../../containers/Cart";
import { Link } from "react-router-dom";
import { MoneyDevise } from "../../../containers/Dropers";

export const NavBarPres = () => (
	<nav>
		<MoneyDevise />
		<div id="nav__routes">
			{navRoutes.map(({ main, route }) => {
				return (
					<Link to={route} key={route}>
						{main}
					</Link>
				);
			})}
		</div>
		<div id="nav__icons">
			{navIcons.map(({ Icon }) => (
				<Icon key={Icon} />
			))}
		</div>
	</nav>
);

const navRoutes: { main: JSX.Element; route: string }[] = [
	{
		main: <p>Home</p>,
		route: "/"
	},
	{
		main: <p>Collections</p>,
		route: "/collection"
	},
	{
		main: <p>News</p>,
		route: "/news"
	}
];

const navIcons: { Icon: any }[] = [
	{
		Icon: LogIcon
	},
	{
		Icon: SearchIcon
	},
	{
		Icon: CartIcon
	}
];

export const NavSpace = () => <div id="nav__space"></div>;
