import React from "react";

export const CarouselPres = () => <></>;

export const SlideShowPres = ({
	images,
	title,
	text,
	link
}: {
	images: string[];
	title?: string;
	text?: string;
	link?: { ref: string; text: string };
}) => (
	<div>
		{true && <button></button>}
		{images.map(img => (
			<img src={img} alt={img} key={img} />
		))}
		{true && <button></button>}
		{title && <h2>{title}</h2>}
		{text && <p>{text}</p>}
		{link && <button>{link.text}</button>}
	</div>
);

export const SlideCommuPres = ({
	images,
	title,
	text,
	link
}: {
	images: string[];
	title?: string;
	text?: string;
	link?: { ref: string; text: string };
}) => (
	<div>
		{true && <button></button>}
		{images.map(img => (
			<img key={img} src={img} alt={img} />
		))}
		{true && <button></button>}
		{title && <h2>{title}</h2>}
		{text && <p>{text}</p>}
		{link && <button>{link.text}</button>}
	</div>
);

export const ThemesPres = ({
	themes
}: {
	themes: {
		title: string;
		link: { ref: string; text: string };
		img: any;
		describe: string;
	}[];
}) => (
	<div>
		{true && <button></button>}
		{themes.map(({ title, link, img, describe }) => (
			<>
				<img key={img} src={img} alt={img} />
				<h2>{title}</h2>
				<p>{describe}</p>
				<button>{link.text}</button>
			</>
		))}
	</div>
);
