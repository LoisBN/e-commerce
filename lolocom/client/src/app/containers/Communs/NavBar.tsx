import React from "react";

import { Link } from "react-router-dom";
import { SearchIcon } from "./Search";
import { NavBarPres } from "../../components/Communs/NavBar";
import { MoneyDevise } from "../Dropers";
import { LogIcon } from "../Login";
import { CartIcon } from "../Cart";

export const NavBar = () => {
	return (
		<NavBarPres>
			<MoneyDevise />
			{navRoutes.map(
				({ main, route }: { main: JSX.Element; route: string }) => {
					return (
						<Link to={route} key={route}>
							{main}
						</Link>
					);
				}
			)}
		</NavBarPres>
	);
};

const navRoutes: { main: JSX.Element; route: string }[] = [
	{
		main: <p>Home</p>,
		route: "/"
	},
	{
		main: <p>Collections</p>,
		route: "/collections"
	},
	{
		main: <p>News</p>,
		route: "/news"
	},
	{
		main: <LogIcon />,
		route: "/login"
	},
	{
		main: <SearchIcon />,
		route: "/search"
	},
	{
		main: <CartIcon />,
		route: "/cart"
	}
];
