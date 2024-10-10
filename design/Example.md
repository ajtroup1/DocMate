# Example Generation

<img src="https://cdn-icons-png.flaticon.com/512/6509/6509613.png" id="main-icon" style="height: 150px;"/>

### Sample description for the project

## Table of Contents
- Project Metadata
1) Package1
    - Files
    - Types
    - Functions
    - Variables
2) Package2
    - Files
    - Types
    - Functions
    - Variables

---
## Package 1
#### *Sample description of the package*
#### Usage of this package
### Dependencies for `Package1`:
- Dependency1 (<a href={Dependency1.Link}>External link</a>)
    - *Description of the dependency*
    - Import via `github.com/sampleuser/sampledependency`
- Dependency2 (<a href={Dependency2.Link}>External link</a>)
    - *Description of the other dependency*
    - Import via `github.com/sampleuser/sampledependency2`

### Types for `Package1`:
- ### `Type1`
    - *Description of the type*
    - Fields:
        - `Field1` (int)
            - *Description of the field*
        - `Field 2` (string)
            - *Description of the field*

### Package-Level Variables for `Package1`:
- ### `Variable1` (int)
    - *Description of the variable*
- ### `Variable2` (string)
    - *Description of the variable*

### Package-Level Functions for `Package1`
- ### `Function1`
    - *Description of the function*
    - Receiver: `Type1`
    - Params:
        - ### `Variable2` (string)
            - *Description of the variable*
    - Return values:
        - (int) *Description of the return value*
        - <p style="color: #ff4949;">(bool) *This is an error return value*</p>
    - HTTP Responses:
        - `404` *Description of error message or remedy*
        - `200` *Description of error message or remedy*
    - Example code usign `Function1`:
        - ```go
            myVar, err := type1.Function1(0, "")
            ```
            - *Description of code snippet and its purpose/usage*

### Files for `Package1`:
- ### `File1.go`
    - *Description of the file*