//  Compare versions correctly (descending order)
const compareVersions = (a: string, b: string): number => {
	const aParts = a.split('.').map(Number);
	const bParts = b.split('.').map(Number);

	for (let i = 0; i < Math.max(aParts.length, bParts.length); i++) {
		const aVal = aParts[i] ?? 0; // Default to 0 if missing
		const bVal = bParts[i] ?? 0;

		if (aVal !== bVal) {
			return aVal - bVal; // Sort in descending order
		}
	}

	return 0;
};

export { compareVersions };
