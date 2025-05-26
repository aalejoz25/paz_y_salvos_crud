package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:CampoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:FormularioPlantillaRespuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemCampoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:ProcesoParametroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/paz_y_salvos_crud/controllers:SemaforoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
