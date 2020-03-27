import React from "react";
import { NewsLetters } from "../../../containers/Communs/Footer/NewsLetter";
import "./index.css";
import { Link } from "react-router-dom";

export const Footer = () => (
	<footer>
		<FooterInfos />
		<FooterOthers />
		<NewsLetters />
	</footer>
);

const FooterInfos = () => (
	<div>
		<h3>Infos</h3>
		<p>
			Lorem ipsum, dolor sit amet consectetur adipisicing elit. Commodi
			iste vitae necessitatibus porro ea cum facere, atque consectetur.
			Enim quo aspernatur possimus sed inventore expedita saepe optio,
			recusandae voluptate qui? Lorem ipsum dolor sit amet consectetur
			adipisicing elit. Eveniet itaque nesciunt eligendi ipsam ipsa
			tempora inventore consequatur ullam magnam adipisci, a, perspiciatis
			dignissimos consectetur, tempore sunt provident non. Aliquid,
			facilis.
		</p>
	</div>
);

const FooterOthers = () => (
	<div>
		<h3>Others</h3>
		{linksFoot.map(({ ref, title }) => (
			<Link key={ref} to={ref}>
				{title}
			</Link>
		))}
	</div>
);

const linksFoot: { ref: string; title: "string" }[] = [];
