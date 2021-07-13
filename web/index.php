<?php
/******************************************************************************

    Front controller - unique entry point of the application.    

    @license    GPL
    @history    2021-07-13 04:45:46+01:00, Thierry Graff : Creation
********************************************************************************/

define('DS', DIRECTORY_SEPARATOR);

// unique global variable of the application - computed in init.php
$G = [];

try{
    
    session_start();
    
    require_once '../src/app/init.php';
    
    require_once '../app/routes.php';
    
    // GO
    require_once $G['CONTROLLER'];
}
catch(Exception $e){
    // TODO graceful
    echo "<br/>EXCEPTION - msg : " . $e->getMessage();
    echo "\n<pre>"; print_r($e->getTraceAsString()); echo "</pre>";
    //echo "\n<pre>"; print_r($e->getTrace()); echo "</pre>";
}

