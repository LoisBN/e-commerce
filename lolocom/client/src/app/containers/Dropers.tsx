import React from "react";

const devises = [{ value: "US $" }, { value: "EU €" }];

export const MoneyDevise = () => {
	return (
		<select>
			{devises.map(({ value }: { value: any }) => (
				<option value={value}>{value}</option>
			))}
		</select>
	);
};
