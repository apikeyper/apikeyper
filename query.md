SQL queries

var aggregateQuery = `SELECT
			date_trunc('minute', created_at) - (date_part('minute', created_at)::integer % 5 || ' minutes')::interval AS interval_start,
			COUNT(*) AS total_usage
		FROM
			api_key_usages
		GROUP BY
			interval_start
		ORDER BY
			interval_start DESC`

var aggregateQueryWithUsage = `SELECT
						date_trunc('minute', created_at) - (date_part('minute', created_at)::integer % 5 || ' minutes')::interval AS interval_start,
						COUNT(*) AS total_usage,
						usage
					FROM
						api_key_usages
					GROUP BY
						interval_start, usage
					ORDER BY
						interval_start DESC`

var aggregateQueryForSuccess = `SELECT
									date_trunc('minute', created_at) - (date_part('minute', created_at)::integer % 5 || ' minutes')::interval AS interval_start,
									COUNT(*) AS total_usage,
									usage
								FROM
									api_key_usages
								WHERE
									usage = 'success'
								GROUP BY
									interval_start, usage
								ORDER BY
									interval_start DESC`