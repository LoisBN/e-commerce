import React from "react";
import { Filters } from "../containers/Filters";
import { Catalogue } from "../components/Catalogue";
import { BarCollection } from "../components/Filters";

export const Collection = () => {
	return (
		<>
			<BarCollection />
			<div id="catalogue_finder">
				<Filters />
				<Catalogue products={[]} />
			</div>
		</>
	);
};
