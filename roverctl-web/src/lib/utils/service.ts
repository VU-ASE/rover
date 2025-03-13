/**
 * Various utilities for working with Fully Qualified Services (as passed around often between roverd and roverctl)
 */

import type { FullyQualifiedService } from '$lib/openapi';

// Two services are equal when:
// - the name is the same
// - the author is the same
// - the version is the same
// nb: this does not incorporate the "as" field
const serviceEqual = (a: FullyQualifiedService, b: FullyQualifiedService) => {
	return a.name === b.name && a.author === b.author && a.version === b.version;
};

// The service identifier is the string that uniquely identifies a service in the shared namespace
// it is the name of the service, unless an "as" field is present, in which case it is the "as" field
const serviceIdentifier = (service: FullyQualifiedService) => {
	return service.as || service.name;
};

// Two services are conflicting when:
// - either the names are the same
// OR
// - the name and the "as" field are the same
// OR
// - the "as" fields are the same
const serviceConflicts = (a: FullyQualifiedService, b: FullyQualifiedService) => {
	return serviceIdentifier(a) === serviceIdentifier(b);
};

export { serviceEqual, serviceIdentifier, serviceConflicts };
