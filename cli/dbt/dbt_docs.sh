echo ""
echo "
 ______   ______   _________  
|_   _  .|_   _   |  _   _  | 
  | |  . \ | |_) ||_/ | | \_| 
  | |  | | |  __ .    | |     
 _| |_.' /_| |__) |  _| |_    
|______.'|_______/  |_____|
_____________________________________________________"
echo ""
echo "#1. Generating docs..."
echo ""
cd /sql_models && dbt docs generate
echo ""
echo "#2. Serving docs..."
echo ""
cd /sql_models && dbt docs serve --port $DBT_DOCS_PORT