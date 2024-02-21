{{ define "sql-counter" }}

select
    count(*) total_items
from
    fb_{{ .e.TableName }}_entities
where
    parent_id is null

{{ end }}

{{ define "sql" }}


WITH RECURSIVE
    fb_template_entities_cte(level, unique_id, {{ join .e.GetSqlFieldNames "," }}) AS (
    select * from (
        SELECT 0, fb_template_entities.unique_id, {{ join .e.GetSqlFields "," }} from fb_template_entities
        where parent_id is null
        (internalCondition)
        limit @limit offset @offset
    )
    UNION ALL
    SELECT fb_template_entities_cte.level+1,fb_template_entities.unique_id, {{ join .e.GetSqlFields "," }}
        FROM fb_template_entities JOIN fb_template_entities_cte ON fb_template_entities.parent_id=fb_template_entities_cte.unique_id
        ORDER BY 2 DESC

    )
SELECT DISTINCT
    fb_template_entities_cte.level,
    fb_template_entities_cte.unique_id,
    {{ join .e.GetSqlFieldNamesAfter "," }}
    
    FROM fb_template_entities_cte

{{ if .e.HasTranslations }}

LEFT JOIN fb_template_entity_polyglots on fb_template_entity_polyglots.linker_id = fb_template_entities_cte.unique_id
and fb_template_entity_polyglots.language_id = '(language)'

{{ end}}

{{ end }}


{{ if eq .type "sql" }}
    {{ template "sql" . }}
{{ else }}
    {{ template "sql-counter" . }}
{{ end  }}