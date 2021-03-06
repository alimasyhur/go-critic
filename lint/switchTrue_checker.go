package lint

import "go/ast"

func init() {
	addChecker(&switchTrueChecker{}, attrSyntaxOnly)
}

type switchTrueChecker struct {
	checkerBase
}

func (c *switchTrueChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects switch-over-bool statements that use explicit `true` tag value"
	d.Before = `
switch true {
case x > y:
}`
	d.After = `
switch {
case x > y:
}`
}

func (c *switchTrueChecker) VisitStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.SwitchStmt); ok {
		if qualifiedName(stmt.Tag) == "true" {
			c.warn(stmt)
		}
	}
}

func (c *switchTrueChecker) warn(cause *ast.SwitchStmt) {
	if cause.Init == nil {
		c.ctx.Warn(cause, "replace 'switch true {}' with 'switch {}'")
	} else {
		c.ctx.Warn(cause, "replace 'switch %s; true {}' with 'switch %s; {}'",
			cause.Init, cause.Init)
	}
}
