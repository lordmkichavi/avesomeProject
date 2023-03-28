package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"sync"
	//"rsc.io/quote"
	//quoteV3 "rsc.io/quote/v3"
	"time"

	iGo "awesomeProject/POO/Interfacs"
	abs "awesomeProject/POO/abstraccion"
	poo "awesomeProject/POO/ejercicio"
	pol "awesomeProject/POO/polimorfismo"
	conc "awesomeProject/general"
	patron "awesomeProject/patrones/objetonulo"
	"awesomeProject/pruebasunitarias"
	rest "awesomeProject/rest/cmd"
	tour "awesomeProject/tour"
	"awesomeProject/tour/paquetes"
)

type curso struct {
	name  string
	price float64
}

type nuevoCurso curso

type nuevoBool bool

const Pi = 3.14
const Truth = true

func (v nuevoBool) String() string {
	if v {
		return "Verdadero"
	} else {
		return "False"
	}
}

func (v nuevoCurso) libre(para string) string {
	fmt.Println(para)
	if v.name == "java" {
		return "cat"
	} else {
		return "dog"
	}
}

func main() {

	//CambiarValores()
	//impresionValoresEstructuras()
	//imprimirValoresAsociadosEstructura()
	//imprimirValoresConPunterosEstructuraas()
	//probarOperacionesConversionTiposDatos()
	//gestionarLlamadasEntreObjetos()
	//probarPersonaloizacionTypes()
	//probarExtensionesConTypes()
	//probarOperacionesConversionTiposDatos()
	//manejarConstantes()
	//manejoCondicionales()
	//manejoCondicionales2()
	//probarLLamadasDefer()
	//probarExtensionesConTypes()
	//gestionarLlamadasEntreObjetos()
	//probarInterfacesSobrecargado()
	//probarInterfacesDesdePaquete()
	//ProbarSwitchCasoEspecial()
	//ProbarFalloCasting()
	//ProbarLLamadoObjetosConAtributos()
	//LlamarPatronMetodoFabrica()
	//ProbarLlamadoPolimorfico()
	//ProbarCambioValoresConPunteros()
	ProbarCambiarValoresPunterosArgumento()
	//MetodoIniciallamadosConcurrencia()
	//MetodoInicialLlamadosRest()
	//InicioServidorRest()
	//EjecutarMetodosMiddleware()
	//pruebaFuncionesTour()
	//ProbarLlamadoPatronObjetoNoNulo()
	//ProbarInicialPruebasUnitarias()

	//fff := "dddd"
	//fmt.Println("cat ", fff)

	//pruebaHoraria()

	//fmt.Println("gato")
	//pruebaVariables()

	//pruebaMap()
	//pruebasConceptuales()

	//probarMultiparte()
	//MetodoInicialLlamadosRest()

	//ProbarConversion()
	//var cadenaFecha = "2022-12-01T12:21:24-04:00"
	//var resFecha = SetDateFormat(cadenaFecha, "2006-01-02")

	//fmt.Println(resFecha)

	//var gg = CortarCadenas("mlc,")

	//fmt.Println(gg)

	//	var a = new([2]string)
	/*var a = make([]string, 0)
	a[0] = "Hello"
	a[1] = "World"*/

	//var contiene = Contains(gg, "mla")

	//fmt.Println(contiene)

	//var hhh string = "TRANSFER_OUT"

	//fmt.Println(strings.EqualFold("TRANSFER_OUT", hhh))

}

type Planet struct {
	name       string
	galaxyname string
}

func (planet Planet) PrependPlanet() {
	planet.name = "Planet " + planet.name
}

func CambiarValores() {
	earth := Planet{name: "Earth", galaxyname: "MilkyWay"}
	earth.PrependPlanet()
	fmt.Println(earth.name)

	jupiter := Planet{name: "Jupiter", galaxyname: "MilkyWay"}
	jupiter.PrependPlanet()
	fmt.Println(jupiter.name)
}

func Contains(arrayStrings []string, chain string) bool {
	var containChain bool
	for _, pivotChain := range arrayStrings {
		if chain == pivotChain {
			containChain = true
		}
	}
	return containChain
}

func CortarCadenas(site string) []string {
	return strings.Split(site, ",")
}

func SetDateFormat(dateString string, dateFormat string) string {
	dateParse, err := time.Parse(dateFormat, dateString)
	dateConvert := ""
	if err != nil {
		fmt.Println(fmt.Sprintf("error when formatting %s the date %s: %v", dateFormat, dateString, err.Error()))
	} else {
		dateConvert = dateParse.Format(dateFormat)
	}
	return dateConvert
}

func ProbarConversion() {
	slice_1 := [3]int{1, 2, 3}
	//var bbb = make([][]string, 2)
	var ccc = [1][len(slice_1)]int{}

	for i, num := range slice_1 {
		ccc[0][i] = num
	}

	var a = [5][1]int{
		{0},
		{4},
		{8},
		{0},
		{4},
	}

	var b = [5][1]int{}

	fmt.Println(b)

	fmt.Printf("t1: %T\n", a)
	fmt.Printf("t1: ", len(a))
	fmt.Println(a)
	fmt.Println(ccc)

	//bbb[0][0] = string(slice_1[0])

	/*ccc[0][0] = "1"
	ccc[0][1] = "2"
	ccc[0][2] = "3"
	ccc[0][3] = "4"
	ccc[0][4] = "5"
	ccc[0][5] = "6"*/

	//fmt.Println(ccc)

	//fmt.Println(bbb)
	//fmt.Println(gato)
}

func pruebasConceptuales() {
	conc.Probar()
}

func uploadFileMultipart(url string, path string) (*http.Response, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Reduce number of syscalls when reading from disk.
	bufferedFileReader := bufio.NewReader(f)
	defer f.Close()

	// Create a pipe for writing from the file and reading to
	// the request concurrently.
	bodyReader, bodyWriter := io.Pipe()
	formWriter := multipart.NewWriter(bodyWriter)

	// Store the first write error in writeErr.
	var (
		writeErr error
		errOnce  sync.Once
	)
	setErr := func(err error) {
		if err != nil {
			errOnce.Do(func() { writeErr = err })
		}
	}
	go func() {
		partWriter, err := formWriter.CreateFormFile("file", path)
		setErr(err)
		_, err = io.Copy(partWriter, bufferedFileReader)
		setErr(err)
		setErr(formWriter.Close())
		setErr(bodyWriter.Close())
	}()

	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", formWriter.FormDataContentType())

	// This operation will block until both the formWriter
	// and bodyWriter have been closed by the goroutine,
	// or in the event of a HTTP error.
	resp, err := http.DefaultClient.Do(req)

	if writeErr != nil {
		return nil, writeErr
	}

	return resp, err
}

func probarMultiparte() {
	path, _ := os.Getwd()
	path += "/Users/jafajardo/GolandProjects/awesomeProject/poder.pdf"
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	request, err := newfileUploadRequest("https://google.com/upload", extraParams, "file", "/Users/jafajardo/GolandProjects/awesomeProject/poder.pdf")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		var bodyContent []byte
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		resp.Body.Read(bodyContent)
		resp.Body.Close()
		fmt.Println(bodyContent)
	}
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}

func pruebaVariables() {

	name := "DEV_RECONCILABLES_DS"
	var writeEndpointVariableFormat = "DOCUMENT_SEARCH_%s_WRITE_ENDPOINT"
	//var resulta = os.Getenv(fmt.Sprintf(writeEndpointVariableFormat, name))
	var resulta = fmt.Sprintf(writeEndpointVariableFormat, name)

	fmt.Println(resulta)
}

func pruebaHoraria() {

	t := time.Now()
	zone, offset := t.Zone()
	fmt.Println(zone, offset)

	fmt.Println(t)
}

func ProbarInicialPruebasUnitarias() {
	//var g = pruebasunitarias.Sumar(1, 2)
	//fmt.Println(g)
	var g = pruebasunitarias.SumarAcomulativa(1, 2, 6, 7)
	fmt.Println(g)
}

func ProbarLlamadoPatronObjetoNoNulo() {
	patron.LlamarImplementacionPatron()
}

func saludarM(nombre string) {
	fmt.Println("saludar ", nombre)
}

func despedirM(nombre string) {
	fmt.Println("despedir ", nombre)
}

func EjecutarMetodosMiddleware() {
	fmt.Println("cat")
	execute("gato", paquetes.MiddlewareLog(saludarM))
	execute("perro", paquetes.MiddlewareLog(despedirM))
}

func execute(nombre string, f paquetes.MyFuncion) {
	f(nombre)
}

func InicioServidorRest() {
	fmt.Println("inicio servidor")
	rest.PruebaCMD()
}

func MetodoInicialLlamadosRest() {
	fmt.Println("car")
	//rest.IniciarServidorWeb()
	rest.PruebaCMD()
}

func MetodoIniciallamadosConcurrencia() {
	//con.SimularProcesoSecuencial()
	//con.SimularProcesoConcurrente()
	//concurrencia.EjecutarMovimientosSecuencial()
	//concurrencia.EjecutarMovimientosConcurrente()
	//concurrencia.PruebaLlamadaGoRutina()
	//concurrencia.PruebaLlamadaGoRutina2()
	//concurrencia.PruebaLlamadaGoRutina3()
	//concurrencia.PruebaCanal1()
	//concurrencia.PruebaCanal2()
	//concurrencia.SimularProcesoConcurrenteCSP()
	//concurrencia.PruebaCanalCancelacionSegura()
}

func ProbarCambiarValoresPunterosArgumento() {
	x := 0
	fmt.Println("x puntero", &x)
	fmt.Println("x valor", x)
	cambiarValores(&x)
	funcionInternaCambioValores(x)
	fmt.Println("x puntero despues", &x)
	fmt.Println("x valor despues", x)
}

func funcionInternaCambioValores(y int) {
	fmt.Println("y antes valor", y)
}

func cambiarValores(y *int) {
	fmt.Println("y antes valor", y)
	fmt.Println("y antes dereferencia", *y)
	*y = 43
	fmt.Printf("Tipo de dato: %T\n", y)
	fmt.Printf("Tipo de dato: %T\n", *y)
	fmt.Printf("Tipo de dato: %T\n", &y)
	fmt.Println("y nuevo valor", y)
	fmt.Println("y nuevo valor de referencia", *y)
}

func ProbarCambioValoresConPunteros() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // 73
}

/*
2701 /
*/

func ProbarLlamadoPolimorfico() {
	pago := pol.MetodoFabrica(1)
	pol.Encapsulador("gato")
	pol.Encapsulador(2)
	pol.Encapsulador(pago)
}

func LlamarPatronMetodoFabrica() {
	pago := pol.MetodoFabrica(1)

	pago.Pagar()
}

func ProbarLLamadoObjetosConAtributos() {
	p1 := iGo.NewAnimal("gato")
	p1.Set("mico")
	var nombre = p1.Get()
	fmt.Println(nombre)

	iGo.Exec(p1, "leon")
	var nombre2 = p1.Get()
	fmt.Println(nombre2)
}

func ProbarFalloCasting() {
	var gggg = 22
	var ccc = string(gggg)
	fmt.Println(ccc)
}

func ProbarSwitchCasoEspecial() {
	nextstop := "B"

	fmt.Println("Stops ahead of us:")

	switch nextstop {

	case "A":
		fmt.Println("A")
		fallthrough

	case "B":
		fmt.Println("B")
		fallthrough

	case "C":
		fmt.Println("C")
		fallthrough

	case "D":
		fmt.Println("D")
		fallthrough

	case "E":
		fmt.Println("E")
	}
}

func probarInterfacesDesdePaquete() {
	p1 := tour.Person{
		Name: "Warren",
		Age:  31,
		Addr: &tour.Address{
			City:    "Denver",
			State:   "CO",
			Country: "U.S.A.",
		},
	}
	fmt.Printf("%#v\n", p1)
	p2 := tour.Person{
		Name: "Theia",
		Age:  4,
	}
	fmt.Printf("%#v\n", p2)
}

func probarInterfacesSobrecargado() {
	var iJav iGo.Multiple = iGo.Persona{
		"javier",
	}
	var iJav2 iGo.Multiple = iGo.Texto("gato")
	//var gg = iJav.Saludar("cat")
	//var bb = iJav2.Saludar("perro")
	//fmt.Println(gg)
	//fmt.Println(bb)
	iGo.SaludarTodos(iJav2, iJav)
	fmt.Println("-------------")
	fmt.Printf("ddde", iJav)
}

func gestionarLlamadasEntreObjetos() {
	cliente := poo.NewCliente("javier", "fajardo", "3185149128")
	producto := poo.NewProducto(1, "galletita", 1.39)
	producto2 := poo.NewProducto(2, "pastelita", 2.89)
	factura := poo.NewFactura(""+
		"colombia",
		"bogota",
		0.0,
		cliente,
		poo.NewProductos(producto, producto2),
	)
	factura.SetTotal()

	empleado := poo.NewEmpleado("ignacio", "coruña", "3185149128", 9000)

	empleado.CalcularNomina()
	var gg = cliente.Saludar("a")

	fmt.Println(gg)

	empleado.Saludar("b")

	fmt.Println(empleado)
	fmt.Println(factura)
}

func probarExtensionesConTypes() {
	curso := curso{
		name:  "javaa",
		price: 345,
	}

	valor := nuevoCurso(curso)

	v := valor.libre("gato")

	fmt.Println(valor)
	fmt.Println(v)
}

func probarLLamadasDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)

	fmt.Println("hello")

	defer funcion1()
	defer funcion2()
	defer funcion3()
}

func funcion1() {
	fmt.Println("a")
}

func funcion2() {
	fmt.Println("b")
}

func funcion3() {
	fmt.Println("c")
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}

	return file.Close()
}

func manejoCondicionales2() {
	/*fmt.Print("Go runs on ")
	os := runtime.GOOS
	fmt.Println("asas", os)
	switch os {
	case "darwisdsdn":
		fmt.Println("OS X.")
	case "linuxff":
		fmt.Println("Linux.")
	}
	fmt.Println(os)*/
	/*today := time.Now().Weekday()
	fmt.Println(today + 0)
	fmt.Println(time.Saturday)
	switch time.Monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}*/
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 1:
		fmt.Println("Good morning!")
	case t.Hour() < 2:
		fmt.Println("Good afternoon.")

	}
}

const DELTA = 0.0000001
const INITIAL_Z = 100.0

func Sqrt(x float64) (z float64) {
	z = INITIAL_Z

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}

	for zz := step(); math.Abs(zz-z) > DELTA; {
		z = zz
		zz = step()
	}
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func manejoCondicionales() {
	/*sum := 1
	for sum < 5 {
		fmt.Println(sum)
		sum = sum + sum
		fmt.Println("asass ", sum)
	}
	fmt.Println(sum)*/

	//fmt.Println(sqrt(-2))

	//var ff = fmt.Sprintln("ssss")
	//fmt.Println(ff)

	/*var x = 123
	if x == 123 {
		fmt.Println("cat")
	} else {
		fmt.Println("dog")
	}*/

	/*var g = pow(2, 3, 9)

	fmt.Println(g)*/

	fmt.Println(Sqrt(3))
}

func pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func manejarConstantes() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//fmt.Println(Big)
	fmt.Println(Small)
}

func probarPersonaloizacionTypes() {
	valor := nuevoBool(true)
	fmt.Println(valor)
}

func probarOperacionesConversionTiposDatos() {
	num1 := uint8(12)
	num2 := uint16(12)
	num3 := int(10)
	num4 := int16(15)
	num5 := int64(45)

	num6 := float64(12.23)
	num7 := float32(23.34)

	num8 := float32(num6) + num7

	fmt.Printf("Tipo de dato: %T\n", num1)
	fmt.Printf("To de dato: %T\n", num2)
	fmt.Printf("Tipe dato: %T\n", num1+uint8(num2))
	fmt.Printf("Tipo de dato: %T\n", int16(num3)+num4)
	fmt.Printf("Tipo de dato: %T\n", num4+int16(num5))
	fmt.Printf("Tipo de dato: %T\n", float32(num6)+num7)
	fmt.Printf("Tipo de dato: %T\n", float64(num3)+num6)
	fmt.Printf("Tipo de dato: %T\n", num8)
	fmt.Printf("Tipo de dato: %T\n", num5+int64(num7))

	var gggg = 22
	var ccc = string(gggg)

	fmt.Printf("Tipo de dato: %T\n", ccc)
	fmt.Printf("struct2: %+v\n", "asasas")

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func imprimirValoresConPunterosEstructuraas() {
	p := abs.New("", 45.67, true)
	p.UserIDs = []uint{2, 3, 4}
	p.Classes = map[uint]string{
		1: "Calculo",
		2: "Fisica",
		3: "Catedra",
	}
	var nuevoPrecio float64 = 23.5
	p.CambiarPrecio(nuevoPrecio)
	p.ImprimirClasses()
	fmt.Println(p.UserIDs)
}

func imprimirValoresAsociadosEstructura() {
	p := abs.New("", 45.67, true)
	p.ImprimirClasses()
	fmt.Println(p.Price)
}

func impresionValoresEstructuras() {
	valoresUno := abs.New("go programming", 45.67, true)
	valoresDos := abs.New("", 45.67, true)
	valoresTres := abs.New("", 45.67, true)
	valoresTres.IsFree = true
	fmt.Printf("%+v\n", valoresUno)
	fmt.Printf("%+v\n", valoresDos)
	fmt.Printf("%+v\n", valoresTres)
}

/*func pruebaDependenciasExyternas() {
	fmt.Println("hello")

	var g = saludo2.SaludarIngles("gatogato")

	var bbb = quoteV3.Concurrency()

	fmt.Println(bbb)
	fmt.Println(g)

	var bb = quote.Hello()

	fmt.Println(bb)
}*/

func incluirVerdadero(numero int) bool {
	return false
}
