package designPatern

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	//TODO implement me
	panic("implement me")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	//TODO implement me
	panic("implement me")
}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

func (o OldFashionedPrinter) Fax(d Document) {
	//TODO implement me
	panic("implement me")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	//TODO implement me
	panic("implement me")
}

var _ Machine = (*OldFashionedPrinter)(nil)
var _ Machine = (*MultiFunctionPrinter)(nil)

// ISP

type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	//TODO implement me
	panic("implement me")
}

func (p Photocopier) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax what ever you want to do here
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}
func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

var _ Scanner = (*Photocopier)(nil)
var _ Printer = (*Photocopier)(nil)
var _ Printer = (*MyPrinter)(nil)
