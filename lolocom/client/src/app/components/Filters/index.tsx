import React from "react";
import c from "./filters.module.css";
import { DisplayModeSwitcherPres } from "../Switchers";
import { DropDownOptions } from "../../containers/Dropers";

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

export const FiltersPrev = ({ children }: { children: any }) => (
	<div id={c.filterscolumn}>{children}</div>
);

export const BarCollection = () => {
	const [o, setO] = React.useState<number>(0);
	const changeOption = (e: any) => {};
	return (
		<div id={c.filtersbar}>
			<DisplayModeSwitcherPres setter={() => {}} current={true} />
			{/*
			<DropDownOptions
				list={listMainOptions}
				current={listMainOptions[o]}
				setter={changeOption}
			/>
			*/}
		</div>
	);
};

const listMainOptions = [
	["Du Plus Grand au plus petit prix", ">"],
	["Du Plus petit au plus Grand prix", "<"]
];
