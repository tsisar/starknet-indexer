{{/*
Common labels
*/}}
{{- define "application.labels" -}}
app.kubernetes.io/name: {{ .Release.Name }}
app.kubernetes.io/instance: {{ .Release.Namespace }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Generate PostgreSQL password if not provided
*/}}
{{- define "postgres.password" -}}
{{- if .Values.postgres.auth.password -}}
{{- .Values.postgres.auth.password -}}
{{- else -}}
{{- randAlphaNum 16 -}}
{{- end -}}
{{- end -}}

{{/*
Generate PostgreSQL username
*/}}
{{- define "postgres.username" -}}
{{- .Values.postgres.auth.username | default "postgres" -}}
{{- end -}}

{{/*
Check if external secrets should be used
*/}}
{{- define "postgres.useExternalSecrets" -}}
{{- if or .Values.postgres.externalSecret.enabled .Values.postgres.secretProviderClass.enabled -}}
{{- true -}}
{{- else -}}
{{- false -}}
{{- end -}}
{{- end -}}
