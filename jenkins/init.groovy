import jenkins.model.*
import org.jenkinsci.plugins.workflow.job.WorkflowJob
import org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition

def jobName = "GoAppPipeline"
def jobScriptPath = "/var/jenkins_home/workspace/GoAppPipeline/Jenkinsfile"

if (Jenkins.instance.getItem(jobName) == null) {
    println("Creating pipeline '${jobName}'...")

    def jenkinsInstance = Jenkins.getInstance()
    def job = jenkinsInstance.createProject(WorkflowJob, jobName)

    def jobScript = new File(jobScriptPath).text
    job.definition = new CpsFlowDefinition(jobScript, true)
    job.save()

    println("Pipeline '${jobName}' created successfully.")
} else {
    println("Pipeline '${jobName}' already exists. Skipping creation.")
}
