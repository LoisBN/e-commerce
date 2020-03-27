import React from "react";

const devises = [{ value: "US $" }, { value: "EU €" }];

export const MoneyDevise = () => {
	return (
		<select>
			{devises.map(({ value }: { value: any }) => (
				<option value={value} key={value}>
					{value}
				</option>
			))}
		</select>
	);
};

export const DropDownOptions = ({
	list,
	current,
	setter
}: {
	list: string[];
	current: string;
	setter: any;
}) => {
	return (
		<div className="dropdown">
			<span>Mouse over me</span>
			<div className="dropdown-content">
				<p>Hello World!</p>
			</div>
		</div>
	);
};

export const NumberChange = () => {
	return (
		<div>
			<button></button>
			<input />
			<button></button>
		</div>
	);
};
