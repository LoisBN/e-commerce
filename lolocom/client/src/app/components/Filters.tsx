import React from "react";

export const MainTile = ({
	title,
	develop,
	choosed,
	isDevelopped
}: {
	title: React.ReactNode;
	develop: any;
	choosed?: React.ReactNode[];
	isDevelopped: boolean;
}) => {
	return (
		<button onClick={develop}>
			{title}
			{choosed}
		</button>
	);
};

export const Tile = ({
	title,
	isChoosed,
	choose
}: {
	title: React.ReactNode;
	isChoosed: boolean;
	choose: any;
}) => {
	return <button onClick={choose}>{title}</button>;
};
