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
echo "Welcome $FLIP_USERNAME! This is an interactive DBT environment
that can be used to run your DBT sql models. 

Try running any 'dbt' command using the prompt below. 

You are currently in the 'sql_models' directory."
echo "
-----------------------------------------------------"
echo "PS1='🕵️‍♀️ [\w]\[\033[00m\] $ '" >> /root/.bashrc
exec /bin/bash