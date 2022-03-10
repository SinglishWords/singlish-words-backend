# Read Password
read -s -p "MySQL Password: " password
# Run Command
mysql -uroot -p$password < sql/initDB.sql < sql/questions-test-data.sql
perl -pi -e 's/csqsiew/root/g' ./config.yaml
perl -pi -e 's/123456/$password/g' ./config.yaml