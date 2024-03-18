import * as ts from "typescript";

// Define a function to visit nodes in the AST (Abstract Syntax Tree)
// and extract the properties of the type alias declaration
const visit = (node: ts.Node, sourceFile: ts.SourceFile): any => {
  const extractProperties = (
    node: ts.TypeNode | undefined
  ): Record<string, any> | Array<any> | string | undefined => {
    // If the node is undefined, return undefined
    if (!node) return;

    switch (node.kind) {
      // If the node is a type reference, extract the properties of the type
      case ts.SyntaxKind.TypeLiteral: {
        const properties: Record<string, any> = {};
        node.forEachChild((member) => {
          if (ts.isPropertySignature(member)) {
            // Extract property value
            const value = extractProperties(member.type);
            // Get property name and pluralize if value is an array
            let propertyName = member.name.getText(sourceFile);
            if (Array.isArray(value)) {
              propertyName += "s";
            }
            // Recursively extract properties and assign to the properties object
            properties[propertyName] = extractProperties(member.type);
          }
        });
        return properties;
      }
      // If the node is a union type, extract the types of the union
      // and return an array of the types
      case ts.SyntaxKind.UnionType: {
        const unionNode = node as ts.UnionTypeNode;
        return unionNode.types.map((type) => extractProperties(type));
      }
      // If the node is a literal type, return the value of the literal and remove the quotes
      case ts.SyntaxKind.LiteralType: {
        return node
          .getText(sourceFile)
          .substring(1, node.getText(sourceFile).length - 1);
      }
      // Otherwise return the text of the node so that it can be used for future comparisons
      default:
        return node.getText(sourceFile);
    }
  };

  // If the string has an alias declaration return an object with the type alias name as key and extracted properties as value
  if (ts.isTypeAliasDeclaration(node)) {
    return {
      [node.name.text.toLowerCase()]: extractProperties(node.type),
    };
  }
};

const convertToObject = (type: string) => {
  // Create a source file in memory from the type string
  const sourceFile = ts.createSourceFile(
    "source.ts",
    type,
    ts.ScriptTarget.ES2020
  );

  // Traverse the AST so visit can look for the type alias declaration and extract properties
  return ts.forEachChild(sourceFile, (node: ts.Node) =>
    visit(node, sourceFile)
  );
};

// Convert the type string to an object and log the result
const object = convertToObject(`type Button = {
  variant: 'solid' | 'text';
};`);

console.log(JSON.stringify(object, null, 2));
